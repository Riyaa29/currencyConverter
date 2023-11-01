<template>
  <div class="container">
    
    <!-- Input for fromCurrency -->
    <div class="input-group">
      <label for="fromCurrency">From Currency:</label>
      <select v-model="fromCurrency" id="fromCurrency">
        <option v-for="currency in currencies" :key="currency" :value="currency">
            {{ currency }}
        </option>
      </select>
    </div>
    
    <!-- Input for toCurrency -->
    <div class="input-group">
      <label for="toCurrency">To Currency:</label>
      <select v-model="toCurrency" id="toCurrency">
        <option v-for="currency in currencies" :key="currency" :value="currency">
            {{ currency }}
        </option>
      </select>
    </div>
    
    <!-- Input for amount -->
    <div class="input-group">
      <label for="amount">Amount:</label>
      <input v-model.number="amount" id="amount" type="number" placeholder="1">
    </div>

    <!-- Convert button -->
    <button @click="convert">Convert</button>
    
    <!-- Display the result -->
    <div class="result">
      <p><strong>Result:</strong> {{ result }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      fromCurrency: 'USD',
      toCurrency: 'EUR',
      amount: 1,
      result: 0,
      currencies: [], // New property
    };
  },
  methods: {
    async convert() {
      try {
        const response = await axios.get('http://localhost:8081/convert', {
          params: {
            from: this.fromCurrency,
            to: this.toCurrency,
            amount: this.amount,
          },
        });
        this.result = response.data.result;
      } catch (error) {
        console.error("Error converting currency:", error);
      }
    },
  },
  mounted() {
    axios.get('http://localhost:8081/currencies')
        .then(response => {
            this.currencies = response.data.currencies;
        })
        .catch(error => {
            console.error("Error fetching currencies:", error);
        });
  },
};
</script>

<style scoped>
/* Container styling */
.container {
  font-family: 'Arial', sans-serif;
  max-width: 400px;
  margin: 50px auto;
  padding: 20px;
  border: 1px solid #e1e1e1;
  border-radius: 5px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  background-color: #f4f4f2; /* Off-white background */
}

/* Logo styling */
.logo {
  display: block;
  max-width: 100px; /* Adjust as needed */
  margin: 0 auto 20px;
}

/* Input group styling */
.input-group {
  margin-bottom: 20px;
}

/* Input, label, and select styling */
label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
  color: #333;
}

input, select {
  width: 100%;
  padding: 10px;
  margin-bottom: 15px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 16px;
}

/* Button styling */
button {
  display: block;
  width: 100%;
  padding: 10px;
  background-color: #007BFF;
  color: #fff;
  border: none;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button:hover {
  background-color: #0056b3;
}

/* Result styling */
.result {
  margin-top: 20px;
}

p {
  font-size: 18px;
  color: #333;
}

strong {
  color: #007BFF;
}
</style>
