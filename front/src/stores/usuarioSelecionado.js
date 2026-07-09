import { defineStore } from 'pinia'

export const useUsuarioSelecionadoStore = defineStore('usuarioSelecionado', {
  state: () => ({
    usuario: null,
  }),
  actions: {
    setUsuario(usuario) {
      this.usuario = usuario
    },
    limparUsuario() {
      this.usuario = null
    },
  },
})
