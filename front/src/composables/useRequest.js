export function useRequest() {
  const request = async (url, method = 'GET', data = null, headers = {}) => {
    try {
      const config = {
        method,
        headers: {
          'Content-Type': 'application/json',
          ...headers,
        },
      }

      if (data && method !== 'GET') {
        config.body = JSON.stringify(data)
      }

      const response = await fetch(url, config)

      const contentType = response.headers.get('content-type')
      const isJson = contentType && contentType.includes('application/json')
      const body = isJson ? await response.json() : await response.text()

      if (!response.ok) {
        throw new Error(body?.erro || body?.message || 'Erro na requisição')
      }

      return body
    } catch (err) {
      throw err
    }
  }

  return { request }
}
