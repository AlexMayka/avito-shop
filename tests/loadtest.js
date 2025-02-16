import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    stages: [
        { duration: '30s', target: 20 }, // Разогрев: увеличение до 20 пользователей за 30 сек
        { duration: '1m', target: 20 },  // Нагрузка: удержание 20 пользователей в течение 1 минуты
        { duration: '30s', target: 0 },  // Плавное снижение до 0
    ],
    thresholds: {
        'http_req_duration': ['p(95)<500'], // 95% запросов должны быть быстрее 500ms
        'http_req_failed': ['rate<0.1'],    // Менее 10% ошибок
    },
};

const BASE_URL = 'http://localhost:8080';
let token = '';

export function setup() {
    const loginRes = http.post(`${BASE_URL}/api/auth`, JSON.stringify({
        username: 'testUser2',
        password: 'testpass'
    }), {
        headers: { 'Content-Type': 'application/json' },
    });

    if (loginRes.status !== 200) {
        console.error('Login failed! Status:', loginRes.status);
        return {};
    }

    try {
        const parsedBody = JSON.parse(loginRes.body);
        if (!parsedBody.token) {
            console.error("No token in response!");
            return {};
        }
        return { token: parsedBody.token };
    } catch (e) {
        console.error('Error parsing login response:', e);
        return {};
    }
}

export default function(data) {
    const headers = {
        'Authorization': `Bearer ${data.token}`,
        'Content-Type': 'application/json',
    };

    // Тестируем получение информации
    const infoRes = http.get(`${BASE_URL}/api/info`, { headers });
    check(infoRes, {
        'info successful': (r) => r.status === 200,
    });

    // Тестируем покупку предмета
    const buyRes = http.get(`${BASE_URL}/api/buy/t-shirt`, { headers });
    check(buyRes, {
        'buy successful': (r) => r.status === 200,
    });

    // Тестируем отправку монет
    const sendRes = http.post(`${BASE_URL}/api/sendCoin`, JSON.stringify({
        toUser: 'testUser',
        amount: 10
    }), { headers });
    check(sendRes, {
        'send successful': (r) => r.status === 200,
    });

    sleep(1);
}