import request from '@/utils/request'

export function getEngineInfo() {
  return request({
    url: '/admin/engine',
    method: 'get'
  })
}

export function UpdateEngineGroups(data) {
  return request({
    url: '/admin/engine',
    method: 'put',
    data: JSON.stringify(data),
    headers: {
      'Content-Type': 'application/json'
    }
  })
}
