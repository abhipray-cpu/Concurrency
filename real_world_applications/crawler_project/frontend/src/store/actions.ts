import Cookies from 'js-cookie'
import axios from 'axios'
const API_BASE = 'http://localhost:8081'
interface Response {
  status: number
  message: string
  data: any
}

interface LoginPayload {
  email: string
  password: string
}

interface SignupPayload {
  username: string
  email: string
  password: string
}
export default {
  checkAuth(): boolean {
    // Retrieve the token from cookies using js-cookie
    const token = Cookies.get('authToken')

    // Check if the token exists
    return !!token // This will return true if the token exists, false otherwise
  },

  async Login(payload: LoginPayload): Promise<Response> {
    try {
      const body = {
        Email: payload.email,
        Password: payload.password
      }
      const reqUrl = `${API_BASE}/login`
      const response = await axios.post(reqUrl, body)
      if (response.status === 200) {
        const token = response.data.token
        Cookies.set('authToken', token)
      }
      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },

  async Signup(payload: SignupPayload): Promise<Response> {
    try {
      const body = {
        Username: payload.username,
        Email: payload.email,
        Password: payload.password
      }
      const reqUrl = `${API_BASE}/signup`

      const response = await axios.post(reqUrl, body)

      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },
  async Profile(): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/account/get`
      const headers = {
        Authorization: `Bearer ${token}`
      }

      const response = await axios.get(reqUrl, {
        headers: headers
      })

      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },
  async Edit(payload: SignupPayload): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const body = {
        Username: payload.username,
        Email: payload.email,
        Password: payload.password
      }

      const reqUrl = `${API_BASE}/account/edit`
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }

      const response = await axios.put(reqUrl, body, {
        headers: headers
      })

      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },

  async Delete(): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/account/delete`
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }

      const response = await axios.delete(reqUrl, {
        headers: headers
      })

      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },

  async GetPages(): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/page/`
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }
      const response = await axios.get(reqUrl, {
        headers: headers
      })

      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },
  async GetPage(id: string): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/page/${id}`
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }

      const response = await axios.get(reqUrl, {
        headers: headers
      })
      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },
  async EditPage(id: string, content: string): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/page/edit/${id}`
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }
      const body = {
        Content: content
      }
      const response = await axios.put(reqUrl, body, {
        headers: headers
      })
      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },
  async DeletePage(id: string): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/page/delete/${id}`
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }

      const response = await axios.delete(reqUrl, {
        headers: headers
      })
      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },
  async AddURL(url: string): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/page/add`
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }
      const body = {
        URL: url
      }

      const response = await axios.post(reqUrl, body, {
        headers: headers
      })
      console.log(response)
      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  },
  async SearchPage(query: string): Promise<Response> {
    try {
      const token = Cookies.get('authToken')
      if (!token) {
        return {
          status: 401,
          message: 'Unauthorized',
          data: ''
        }
      }
      const reqUrl = `${API_BASE}/page/search`
      const body = {
        Query: query
      }
      const headers = {
        Authorization: `Bearer ${Cookies.get('authToken')}`
      }

      const response = await axios.post(reqUrl, body, {
        headers: headers
      })
      return {
        status: response.status,
        message: response.statusText,
        data: response.data
      }
    } catch (err: any) {
      return {
        status: err.response.status,
        message: err.response.statusText,
        data: ''
      }
    }
  }
}
