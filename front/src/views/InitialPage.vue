<template>
  <div class="home-page">
    <!-- Floating Books Animation -->
    <div class="floating-books">
      <div 
        class="book" 
        v-for="n in 9" 
        :key="n"
        :style="{ left: `${n * 10}%`, animationDelay: `${n * 0.5}s` }"
      ></div>
    </div>

    <!-- Header -->
    <header class="header">
      <div class="container">
        <nav class="navbar">
          <router-link to="/" class="logo">
            <div class="logo-icon">📚</div>
            LivroLog
          </router-link>
          <div class="auth-buttons">
            <button class="btn btn-outline" @click="goToLogin">Login</button>
            <button class="btn btn-primary" @click="goToRegister">Cadastrar</button>
          </div>
        </nav>
      </div>
    </header>

    <!-- Hero Section -->
    <main class="hero">
      <div class="container">
        <div class="hero-content">
          <h1 class="hero-title">Descubra sua próxima leitura</h1>
          <p class="hero-subtitle">
            A plataforma definitiva para avaliar livros, compartilhar reviews e descobrir se aquele livro vale mesmo a pena ser lido.
          </p>

          <!-- Features Section -->
          <div class="features">
            <div class="feature-card" v-for="feature in features" :key="feature.id">
              <div class="feature-icon">{{ feature.icon }}</div>
              <h3 class="feature-title">{{ feature.title }}</h3>
              <p class="feature-description">{{ feature.description }}</p>
            </div>
          </div>

          <!-- Stats Section -->
          <div class="stats">
            <div class="stat-item" v-for="stat in stats" :key="stat.label">
              <span class="stat-number">{{ stat.animatedValue }}</span>
              <span class="stat-label">{{ stat.label }}</span>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'

export default {
  name: 'HomePage',
  setup() {
    const router = useRouter()

    // Reactive data
    const features = ref([
      {
        id: 1,
        icon: '⭐',
        title: 'Avaliações Confiáveis',
        description: 'Leia reviews honestos e detalhados de outros leitores para tomar a melhor decisão antes de investir seu tempo.'
      },
      {
        id: 2,
        icon: '📖',
        title: 'Biblioteca Infinita',
        description: 'Explore milhares de livros de todos os gêneros e descubra obras que você nunca imaginou que existiam.'
      },
      {
        id: 3,
        icon: '👥',
        title: 'Comunidade Ativa',
        description: 'Conecte-se com outros apaixonados por leitura e compartilhe suas descobertas literárias.'
      }
    ])

    const stats = ref([
      { label: 'Livros Avaliados', finalValue: 15000, animatedValue: 0 },
      { label: 'Reviews Publicados', finalValue: 47000, animatedValue: 0 },
      { label: 'Usuários Ativos', finalValue: 12000, animatedValue: 0 },
      { label: 'Avaliação Média', finalValue: 4.8, animatedValue: 0, isDecimal: true }
    ])

    // Navigation methods
    const goToLogin = () => {
      router.push('/login')
    }

    const goToRegister = () => {
      router.push('/register')
    }

    // Animation method
    const animateStats = () => {
      const duration = 2000
      const start = Date.now()
      
      const animate = () => {
        const elapsed = Date.now() - start
        const progress = Math.min(elapsed / duration, 1)
        
        stats.value.forEach(stat => {
          if (stat.isDecimal) {
            stat.animatedValue = (stat.finalValue * progress).toFixed(1)
          } else {
            stat.animatedValue = Math.floor(stat.finalValue * progress)
          }
        })
        
        if (progress < 1) {
          requestAnimationFrame(animate)
        }
      }
      
      requestAnimationFrame(animate)
    }

    // Lifecycle
    onMounted(() => {
      animateStats()
    })

    return {
      features,
      stats,
      goToLogin,
      goToRegister
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

.home-page {
  height: 100vh;
  background: #1a1a1a;
  margin: 0;
  padding: 0;
  width: 100%;
  overflow: hidden;
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 20px;
  width: 100%;
}

.header {
  padding: 15px 0;
  position: relative;
  z-index: 100;
  width: 100%;
}

.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: rgba(44, 62, 80, 0.3);
  backdrop-filter: blur(10px);
  border-radius: 20px;
  padding: 15px 30px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
}

.navbar:hover {
  background: rgba(44, 62, 80, 0.5);
  transform: translateY(-2px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
}

.logo {
  font-size: 28px;
  font-weight: 800;
  color: white;
  text-decoration: none;
  display: flex;
  align-items: center;
  gap: 10px;
}

.logo-icon {
  width: 40px;
  height: 40px;
  background: #2c3e50;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  animation: pulse 2s infinite;
  box-shadow: 0 4px 15px rgba(44, 62, 80, 0.5);
}

@keyframes pulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.05); }
}

.auth-buttons {
  display: flex;
  gap: 15px;
  align-items: center;
}

.btn {
  padding: 12px 24px;
  border: none;
  border-radius: 25px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 14px;
  position: relative;
  overflow: hidden;
}

.btn::before {
  content: '';
  position: absolute;
  top: 0;
  left: -100%;
  width: 100%;
  height: 100%;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
  transition: left 0.5s;
}

.btn:hover::before {
  left: 100%;
}

.btn-outline {
  background: transparent;
  color: white;
  border: 2px solid rgba(255, 255, 255, 0.3);
}

.btn-outline:hover {
  background: rgba(255, 255, 255, 0.1);
  border-color: white;
  transform: translateY(-2px);
}

.btn-primary {
  background: #2c3e50;
  color: white;
  border: none;
}

.btn-primary:hover {
  background: #34495e;
  transform: translateY(-3px);
  box-shadow: 0 10px 30px rgba(44, 62, 80, 0.4);
}

.hero {
  padding: 30px 0;
  text-align: center;
  position: relative;
  width: 100%;
  z-index: 50;
  height: calc(100vh - 70px);
  display: flex;
  align-items: center;
}

.hero::before {
  content: '';
  position: absolute;
  top: -50%;
  left: -50%;
  width: 200%;
  height: 200%;
  background: radial-gradient(circle, rgba(44, 62, 80, 0.1) 0%, transparent 70%);
  animation: rotate 20s linear infinite;
  pointer-events: none;
}

@keyframes rotate {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.hero-content {
  position: relative;
  z-index: 10;
}

.hero-title {
  font-size: clamp(2.5rem, 6vw, 3.5rem);
  font-weight: 900;
  color: white;
  margin-bottom: 15px;
  line-height: 1.2;
  text-shadow: 0 4px 20px rgba(0, 0, 0, 0.5);
}

.hero-subtitle {
  font-size: clamp(1rem, 3vw, 1.2rem);
  color: rgba(255, 255, 255, 0.8);
  margin-bottom: 25px;
  max-width: 600px;
  margin-left: auto;
  margin-right: auto;
  line-height: 1.6;
}

.hero-buttons {
  display: flex;
  gap: 20px;
  justify-content: center;
  flex-wrap: wrap;
  margin-bottom: 30px;
}

.hero-buttons .btn {
  padding: 16px 32px;
  font-size: 16px;
  min-width: 160px;
}

.features {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 20px;
  margin-top: 30px;
}

.feature-card {
  background: rgba(44, 62, 80, 0.2);
  backdrop-filter: blur(10px);
  border-radius: 15px;
  padding: 20px;
  text-align: center;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
  position: relative;
  overflow: hidden;
}

.feature-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 3px;
  background: #2c3e50;
  transform: scaleX(0);
  transform-origin: left;
  transition: transform 0.3s ease;
}

.feature-card:hover::before {
  transform: scaleX(1);
}

.feature-card:hover {
  transform: translateY(-10px);
  background: rgba(44, 62, 80, 0.3);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.3);
}

.feature-icon {
  width: 50px;
  height: 50px;
  background: #2c3e50;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0 auto 15px;
  font-size: 20px;
  color: white;
  box-shadow: 0 4px 15px rgba(44, 62, 80, 0.5);
}

.feature-title {
  font-size: 18px;
  font-weight: 700;
  color: white;
  margin-bottom: 10px;
}

.feature-description {
  color: rgba(255, 255, 255, 0.7);
  line-height: 1.4;
  font-size: 14px;
}

.stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
  gap: 20px;
  margin-top: 40px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  background: rgba(44, 62, 80, 0.2);
  border-radius: 15px;
  border: 1px solid rgba(255, 255, 255, 0.1);
  transition: all 0.3s ease;
}

.stat-item:hover {
  background: rgba(44, 62, 80, 0.3);
  transform: translateY(-5px);
}

.stat-number {
  font-size: 2rem;
  font-weight: 900;
  color: white;
  margin-bottom: 5px;
}

.stat-label {
  font-size: 0.9rem;
  color: rgba(255, 255, 255, 0.7);
  text-align: center;
}

.floating-books {
  position: fixed;
  width: 100vw;
  height: 100vh;
  pointer-events: none;
  overflow: hidden;
  top: 0;
  left: 0;
  z-index: 1;
}

.book {
  position: absolute;
  width: 30px;
  height: 40px;
  background: rgba(44, 62, 80, 0.3);
  border: 2px solid rgba(255, 255, 255, 0.1);
  border-radius: 3px;
  opacity: 0.5;
  animation: float 6s ease-in-out infinite;
}

@keyframes float {
  0%, 100% { transform: translateY(100vh) rotate(0deg); }
  50% { transform: translateY(-100px) rotate(180deg); }
}

@media (max-width: 768px) {
  .auth-buttons {
    flex-direction: column;
    gap: 10px;
  }
  
  .hero-buttons {
    flex-direction: column;
    align-items: center;
  }
  
  .navbar {
    flex-direction: column;
    gap: 20px;
    padding: 20px;
  }
  
  .features {
    grid-template-columns: 1fr;
  }
  
  .stats {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>