import http from 'k6/http';
import { check, sleep, fail } from 'k6';
import { randomString } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

export const options = {
    setupTimeout: '10m', // Увеличенный таймаут до 10 минут
    stages: [
        { duration: '30s', target: 1000 },
        { duration: '1m', target: 1000 },
        { duration: '30s', target: 0 },
    ],
    thresholds: {
        http_req_duration: ['p(95)<50'], // 95% запросов должны быть быстрее 50ms
        http_req_failed: ['rate<0.01'], // Не более 1% ошибок
    },
};

const BASE_URL = 'http://localhost:8080';
const TOTAL_USERS = 100000;
const BATCH_SIZE = 500; // Размер батча

function generateUniqueUser(index) {
    return {
        username: `user_${index}`,
        password: `pass_${index}`
    };
}

export function setup() {
    const users = [];

    for (let batch = 0; batch < TOTAL_USERS / BATCH_SIZE; batch++) {
        console.log(`Создание батча ${batch + 1}/${TOTAL_USERS / BATCH_SIZE}`);

        const batchUsers = Array.from({ length: BATCH_SIZE }, (_, i) => generateUniqueUser(batch * BATCH_SIZE + i));

        const responses = http.batch(
            batchUsers.map(user => ({
                method: 'POST',
                url: `${BASE_URL}/api/auth`,
                body: JSON.stringify(user),
                params: { headers: { 'Content-Type': 'application/json' } }
            }))
        );

        responses.forEach((res, index) => {
            if (res.status === 200) {
                try {
                    const body = JSON.parse(res.body);
                    if (!body.token) throw new Error("Missing token");
                    users.push({ ...batchUsers[index], token: body.token });
                } catch (e) {
                    console.error(`Ошибка обработки ответа: ${e.message}`);
                }
            } else {
                console.error(`Ошибка аутентификации ${batchUsers[index].username}: ${res.status}`);
            }
        });

        console.log(`Создано пользователей: ${users.length}/${TOTAL_USERS}`);
    }

    if (users.length < TOTAL_USERS) {
        throw new Error(`Создано только ${users.length} из ${TOTAL_USERS} пользователей`);
    }

    console.log('Все пользователи успешно зарегистрированы.');
    return { users };
}

export default function (data) {
    const sender = data.users[Math.floor(Math.random() * data.users.length)];

    const headers = {
        Authorization: `Bearer ${sender.token}`,
        'Content-Type': 'application/json',
    };

    const infoRes = http.get(`${BASE_URL}/api/info`, { headers });
    check(infoRes, {
        'info successful': (r) => r.status === 200 || fail(`Ошибка /api/info: ${r.status}`),
    });

    const buyRes = http.get(`${BASE_URL}/api/buy/t-shirt`, { headers });
    check(buyRes, {
        'buy successful': (r) => r.status === 200 || fail(`Ошибка /api/buy/t-shirt: ${r.status}`),
    });

    const recipient = data.users.find(u => u.username !== sender.username);
    if (recipient) {
        const sendRes = http.post(`${BASE_URL}/api/sendCoin`, JSON.stringify({
            toUser: recipient.username,
            amount: 10
        }), { headers });

        check(sendRes, {
            'send successful': (r) => r.status === 200 || fail(`Ошибка /api/sendCoin: ${r.status}`),
        });
    }

    sleep(1);
}

export function teardown(data) {
    console.log(`Тест завершён. Всего создано пользователей: ${data.users.length}`);
}