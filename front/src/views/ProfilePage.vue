<template>
  <div class="profile-page">
    <!-- Header Component -->
    <AppHeader />

    <!-- Main Content -->
    <main class="main-content">
      <!-- Conteúdo da página de perfil -->
      <div class="content-container">
        <!-- Componente de Informações do Perfil -->
        <InfoPerfil :user="userProfile" @profile-updated="handleProfileUpdate" />
      </div>
    </main>
  </div>
</template>

<script>
import AppHeader from '@/components/AppHeader.vue'
import InfoPerfil from '@/components/InfoPerfil.vue'
import { useRequest } from '@/composables/useRequest'

const { request } = useRequest()
export default {
  name: 'ProfilePage',
  components: {
    AppHeader,
    InfoPerfil
  },
  data() {
    return {
      // Dados de exemplo do perfil do usuário
      userProfile: {
        name: 'João Silva',
        email: 'joao.silva@email.com',
        avatar: '/images/avatars/joao.jpg',
        bio: 'Apaixonado por livros de ficção científica e fantasia. Sempre em busca de novas aventuras literárias!',
        reviewsCount: 42,
        averageRating: 4.2,
        memberSince: 'Jan 2023',
        recentReviews: [
          {
            id: 1,
            bookTitle: 'O Senhor dos Anéis: A Sociedade do Anel',
            rating: 5,
            comment: 'Obra-prima da literatura fantástica! Tolkien criou um mundo incrível.',
            date: '2024-06-15T10:30:00Z'
          },
          {
            id: 2,
            bookTitle: 'Duna',
            rating: 4,
            comment: 'Ficção científica complexa e envolvente. Herbert construiu um universo fascinante.',
            date: '2024-06-10T14:22:00Z'
          },
          {
            id: 3,
            bookTitle: '1984',
            rating: 5,
            comment: 'Distopia assustadoramente atual. Orwell foi um visionário.',
            date: '2024-06-05T09:15:00Z'
          },
          {
            id: 4,
            bookTitle: 'Harry Potter e a Pedra Filosofal',
            rating: 4,
            comment: null, // Review apenas com nota
            date: '2024-05-28T16:45:00Z'
          },
          {
            id: 5,
            bookTitle: 'O Hobbit',
            rating: 5,
            comment: 'Perfeito para começar no mundo de Tolkien. Aventura encantadora!',
            date: '2024-05-20T11:20:00Z'
          }
        ]
      }
    }
  },
  methods: {
    handleProfileUpdate(updatedProfile) {
      // Atualizar o perfil com os novos dados
      this.userProfile = { ...this.userProfile, ...updatedProfile }

      // Aqui você pode fazer uma chamada para a API para salvar no backend
      console.log('Perfil atualizado:', updatedProfile)

      // Exemplo de como você poderia fazer uma chamada para API:
      // this.updateProfileInAPI(updatedProfile)

      // Mostrar mensagem de sucesso
      this.showSuccessMessage()
    },

    showSuccessMessage() {
      // Aqui você pode implementar uma notificação de sucesso
      alert('Perfil atualizado com sucesso!')
    },
    async buscarPerfilUsuario(usuarioId) {
      try {
        const url = `http://localhost:8080/usuarios/${usuarioId}/detalhes`
        const dados = await request(url, 'GET')
        this.userProfile = dados
        console.log('Perfil carregado:', dados)
      } catch (e) {
        this.erro = 'Erro ao buscar perfil'
        console.error(e)
      }
    },
    async atualizarPerfilUsuario() {
      try {
        const usuarioId = this.userProfile.id
        const url = `http://localhost:8080/usuarios/${usuarioId}`

        // Monta o payload
        const payload = {
          nome: this.userProfile.nome,
          email: this.userProfile.email,
          descricao: this.userProfile.descricao
        }

        const resultado = await request(url, 'PUT', payload)
        console.log('Perfil atualizado:', resultado)
        alert('Perfil atualizado com sucesso!')
      } catch (e) {
        this.erro = 'Erro ao atualizar perfil'
        console.error(e)
      }
    },
    async deletarPerfilUsuario() {
      const confirmacao = confirm('Tem certeza que deseja deletar seu perfil? Essa ação não pode ser desfeita.')
      if (!confirmacao) return

      try {
        const usuarioId = this.userProfile.id
        const url = `http://localhost:8080/usuarios/${usuarioId}`

        await request(url, 'DELETE')
        alert('Perfil deletado com sucesso!')

        // Redireciona ou limpa dados
        this.userProfile = null
        // this.$router.push('/login') — se estiver usando router
      } catch (e) {
        this.erro = 'Erro ao deletar perfil'
        console.error(e)
      }
    },
    mounted() {
      const usuarioId = localStorage.getItem('usuario_id') // ou outra forma
      if (usuarioId) {
        this.buscarPerfilUsuario(usuarioId)
      }
    }



  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.profile-page {
  min-height: 100vh;
  background-color: #1a1a1a;
  /* Fundo escuro */
  margin: 0;
  padding: 0;
  width: 100%;
  overflow-x: hidden;
}

.main-content {
  padding-top: 64px;
  /* Altura do header fixo */
  height: calc(100vh - 64px);
  /* Preenche o resto da tela */
  width: 100%;
  background-color: #1a1a1a;
  /* Fundo escuro que contrasta com o header #2c3e50 */
  margin: 0;
  padding-left: 0;
  padding-right: 0;
  position: absolute;
  left: 0;
  right: 0;
  top: 64px;
  overflow-x: hidden;
  overflow-y: auto;
  /* Permite scroll vertical */
}

/* Container principal do conteúdo */
.content-container {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: stretch;
  padding: 40px 20px;
  width: 100%;
  min-height: 100%;
  max-width: 1200px;
  margin: 0 auto;
}

/* Responsividade */
@media (max-width: 768px) {
  .content-container {
    padding: 30px 15px;
  }
}

@media (max-width: 480px) {
  .content-container {
    padding: 20px 10px;
  }
}
</style>