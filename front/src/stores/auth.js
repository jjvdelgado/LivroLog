import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    usuario: null, // { id, nome, email, ... }
  }),
  actions: {
    setUsuario(usuario) {
      this.usuario = usuario
      localStorage.setItem('usuario', JSON.stringify(usuario))
    },
    carregarUsuarioLocalStorage() {
      const dados = localStorage.getItem('usuario')
      if (dados) {
        this.usuario = JSON.parse(dados)
      }
    },
    logout() {
      this.usuario = null
      localStorage.removeItem('usuario')
    },
  },
  getters: {
    estaLogado: (state) => !!state.usuario,
  },
})
