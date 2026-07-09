import { defineStore } from 'pinia'

export const useLivroStore = defineStore('livro', {
  state: () => ({
    livroSelecionado: null,
  }),
  actions: {
    setLivro(livro) {
      this.livroSelecionado = livro
    },
    limparLivro() {
      this.livroSelecionado = null
    },
  },
})
