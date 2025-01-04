// Select elements
const expenseForm = document.getElementById('expense-form');
const expenseTableBody = document.querySelector('#expense-table tbody');
const expenseSummary = document.getElementById('expense-summary');
const expenseChart = document.getElementById('expense-chart');

let expenses = JSON.parse(localStorage.getItem('expenses')) || [];

// Function to render expenses
function renderExpenses() {
  // Clear existing table rows and summary
  expenseTableBody.innerHTML = '';
  expenseSummary.innerHTML = '';

  const categoryTotals = {};

  // Add rows for each expense
  expenses.forEach((expense, index) => {
    const row = document.createElement('tr');
    row.innerHTML = `
      <td>₹${expense.amount}</td>
      <td>${expense.description}</td>
      <td>${expense.category}</td>
      <td>
        <button class="delete" data-index="${index}">Delete</button>
      </td>
    `;
    expenseTableBody.appendChild(row);

    // Calculate category totals
    if (!categoryTotals[expense.category]) {
      categoryTotals[expense.category] = 0;
    }
    categoryTotals[expense.category] += expense.amount;
  });

  // Display category totals
  for (const category in categoryTotals) {
    const li = document.createElement('li');
    li.textContent = `${category}: ₹${categoryTotals[category]}`;
    expenseSummary.appendChild(li);
  }

  // Render chart
  renderChart(categoryTotals);
}

// Function to render the chart
function renderChart(categoryTotals) {
  const labels = Object.keys(categoryTotals);
  const data = Object.values(categoryTotals);

  new Chart(expenseChart, {
    type: 'pie',
    data: {
      labels: labels,
      datasets: [{
        label: 'Spending by Category',
        data: data,
        backgroundColor: ['#007bff', '#ffc107', '#28a745', '#dc3545'],
        borderColor: ['#fff'],
        borderWidth: 1,
      }]
    }
  });
}

// Add event listener to the form
expenseForm.addEventListener('submit', (e) => {
  e.preventDefault();

  const amount = parseFloat(document.getElementById('amount').value);
  const description = document.getElementById('description').value;
  const category = document.getElementById('category').value;

  if (amount && description && category) {
    expenses.push({ amount, description, category });
    localStorage.setItem('expenses', JSON.stringify(expenses));
    expenseForm.reset();
    renderExpenses();
  }
});

// Add event listener to delete buttons
expenseTableBody.addEventListener('click', (e) => {
  if (e.target.classList.contains('delete')) {
    const index = e.target.dataset.index;
    expenses.splice(index, 1);
    localStorage.setItem('expenses', JSON.stringify(expenses));
    expenseForm.reset();
    renderExpenses();
  }
});