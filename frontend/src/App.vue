<template>
  <div class="app-container">
    <div class="container">
      <h1>Password Generator</h1>
      
      <div class="password-display">
        <input 
          type="text" 
          :value="password" 
          readonly 
          ref="passwordInput"
          :placeholder="password ? '' : 'Your password will appear here'"
        />
        <button @click="copyToClipboard" :class="{ copied }">
          <span v-if="!copied">Copy</span>
          <span v-else>✓ Copied!</span>
        </button>
      </div>

      <div class="options">
        <div class="length-control">
          <div class="length-header">
            <label>Password Length</label>
            <span class="length-value">{{ length }}</span>
          </div>
          <input 
            type="range" 
            v-model="length" 
            min="4" 
            max="32"
            class="slider" 
          />
          <div class="length-markers">
            <span>4</span>
            <span>18</span>
            <span>32</span>
          </div>
        </div>

        <div class="checkboxes">
          <label class="checkbox-container">
            <input type="checkbox" v-model="uppercase" />
            <div class="checkbox-custom">
              <span class="checkmark">✓</span>
            </div>
            <span class="label-text">Uppercase (A-Z)</span>
          </label>
          
          <label class="checkbox-container">
            <input type="checkbox" v-model="lowercase" />
            <div class="checkbox-custom">
              <span class="checkmark">✓</span>
            </div>
            <span class="label-text">Lowercase (a-z)</span>
          </label>
          
          <label class="checkbox-container">
            <input type="checkbox" v-model="numbers" />
            <div class="checkbox-custom">
              <span class="checkmark">✓</span>
            </div>
            <span class="label-text">Numbers (0-9)</span>
          </label>
          
          <label class="checkbox-container">
            <input type="checkbox" v-model="symbols" />
            <div class="checkbox-custom">
              <span class="checkmark">✓</span>
            </div>
            <span class="label-text">Symbols (!@#$%^&*)</span>
          </label>
        </div>
      </div>

      <button class="generate-btn" @click="generatePassword">
        Generate Password
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const length = ref(12)
const numbers = ref(true)
const symbols = ref(true)
const lowercase = ref(true)
const uppercase = ref(true)
const password = ref('')
const copied = ref(false)
const passwordInput = ref(null)

async function generatePassword() {
  try {
    const apiUrl = import.meta.env.VITE_API_URL || 'http://localhost:8080'
    const response = await fetch(`${apiUrl}/generate`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        length: length.value,
        numbers: numbers.value,
        symbols: symbols.value,
        lowercase: lowercase.value,
        uppercase: uppercase.value,
      }),
    })
    const data = await response.json()
    password.value = data.password
    copied.value = false
  } catch (error) {
    console.error('Error generating password:', error)
  }
}

function copyToClipboard() {
  passwordInput.value.select()
  document.execCommand('copy')
  copied.value = true
  setTimeout(() => {
    copied.value = false
  }, 2000)
}
</script>

<style scoped>
.app-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 1rem;
}

.container {
  width: 100%;
  max-width: 500px;
  background: white;
  padding: 2rem;
  border-radius: 16px;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.15);
}

h1 {
  text-align: center;
  color: #2d3748;
  margin-bottom: 2rem;
  font-size: 1.8rem;
  font-weight: 700;
}

.password-display {
  background: #f7fafc;
  border: 2px solid #e2e8f0;
  border-radius: 8px;
  display: flex;
  margin-bottom: 2rem;
  overflow: hidden;
}

.password-display input {
  flex: 1;
  padding: 1rem;
  font-family: 'Courier New', monospace;
  font-size: 1.1rem;
  border: none;
  background: transparent;
  color: #2d3748;
}

.password-display input::placeholder {
  color: #a0aec0;
}

.password-display button {
  padding: 0.5rem 1.5rem;
  border: none;
  background: #4299e1;
  color: white;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.password-display button:hover {
  background: #3182ce;
}

.password-display button.copied {
  background: #48bb78;
}

.options {
  margin-bottom: 2rem;
}

.length-control {
  margin-bottom: 2rem;
}

.length-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 0.5rem;
}

.length-header label {
  color: #4a5568;
  font-weight: 600;
}

.length-value {
  background: #edf2f7;
  padding: 0.25rem 0.75rem;
  border-radius: 999px;
  font-weight: 600;
  color: #4a5568;
}

.slider {
  -webkit-appearance: none;
  width: 100%;
  height: 6px;
  border-radius: 3px;
  background: #e2e8f0;
  outline: none;
  margin: 1rem 0;
}

.slider::-webkit-slider-thumb {
  -webkit-appearance: none;
  appearance: none;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: #4299e1;
  cursor: pointer;
  transition: all 0.2s;
}

.slider::-webkit-slider-thumb:hover {
  background: #3182ce;
  transform: scale(1.1);
}

.length-markers {
  display: flex;
  justify-content: space-between;
  color: #718096;
  font-size: 0.875rem;
}

.checkboxes {
  display: grid;
  gap: 1rem;
}

.checkbox-container {
  display: flex;
  align-items: center;
  cursor: pointer;
}

.checkbox-container input[type="checkbox"] {
  display: none;
}

.checkbox-custom {
  width: 20px;
  height: 20px;
  border: 2px solid #e2e8f0;
  border-radius: 4px;
  margin-right: 0.75rem;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.2s;
}

.checkbox-container input[type="checkbox"]:checked + .checkbox-custom {
  background: #4299e1;
  border-color: #4299e1;
}

.checkmark {
  color: white;
  font-size: 0.875rem;
  opacity: 0;
  transform: scale(0);
  transition: all 0.2s;
}

.checkbox-container input[type="checkbox"]:checked + .checkbox-custom .checkmark {
  opacity: 1;
  transform: scale(1);
}

.label-text {
  color: #4a5568;
  font-weight: 500;
}

.generate-btn {
  width: 100%;
  padding: 1rem;
  border: none;
  border-radius: 8px;
  background: #4299e1;
  color: white;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}

.generate-btn:hover {
  background: #3182ce;
  transform: translateY(-1px);
}

.generate-btn:active {
  transform: translateY(0);
}

@media (max-width: 480px) {
  .container {
    padding: 1.5rem;
  }

  h1 {
    font-size: 1.5rem;
  }

  .password-display input {
    font-size: 1rem;
    padding: 0.75rem;
  }

  .password-display button {
    padding: 0.5rem 1rem;
  }
}
</style>
