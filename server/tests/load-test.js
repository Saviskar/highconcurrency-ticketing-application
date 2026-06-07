import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 500,
  duration: '2s',
};

export default function () {
  const payload = JSON.stringify({
    slot_id: 1,
    user: `user_${__VU}`,
  });

  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const res = http.post('http://localhost:8080/api/v1/slots/book', payload, params);

  check(res, {
    'status is 200 or 409': (r) => r.status === 200 || r.status === 409,
  });
}
