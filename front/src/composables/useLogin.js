export function useLogin() {
  const login = async (email, senha) => {
    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, senha }),
      })

      if (!response.ok) {
        const erro = await response.json()
        throw new Error(erro.erro || 'Erro no login')
      }

      const data = await response.json()
      return data
    } catch (err) {
      throw err
    }
  }

  return { login }
}
