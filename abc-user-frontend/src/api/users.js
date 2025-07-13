import axios from 'axios'
import { API_BASE_URL } from '@/config'

const client = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export default {
  getUsers(params) {
    return client.get('/users', { params })
  },
  getUser(id) {
    return client.get(`/users/${id}`)
  },
  createUser(data) {
    return client.post('/users', data)
  },
  updateUser(id, data) {
    return client.put(`/users/${id}`, data)
  },
  deleteUser(id) {
    return client.delete(`/users/${id}`)
  },
  checkEmailUnique(email) {
    return client.get('/users/check-email', { params: { email } })
  },
}
