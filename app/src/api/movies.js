import request from '@/utils/request'

export function getMovieLists(req) {
  return request({
    url: '/api/v1/movie',
    method: 'put',
    data: JSON.stringify(req),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}
