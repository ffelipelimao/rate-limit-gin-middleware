import http from 'k6/http'
import { check, sleep } from 'k6'

export default function () {
  let res = http.get('http://localhost:8080/test')

  check(res, { 'ok requests': (r) => r.status === 200 })

  sleep(0.3)
}
