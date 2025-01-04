// Select elements
const taskInput = document.getElementById('task-input');
const addTaskBtn = document.getElementById('add-task-btn');
const taskList = document.getElementById('task-list');
const pendingCount = document.getElementById('pending-count');

// Load tasks from localStorage
let tasks = JSON.parse(localStorage.getItem('tasks')) || [];
renderTasks();

// Event listener to add a task
addTaskBtn.addEventListener('click', () => {
  const taskName = taskInput.value.trim();
  if (taskName) {
    tasks.push({ name: taskName, completed: false });
    taskInput.value = '';
    saveAndRenderTasks();
  }
});

// Function to render tasks
function renderTasks() {
  taskList.innerHTML = '';
  tasks.forEach((task, index) => {
    const taskItem = document.createElement('div');
    taskItem.className = `task-item ${task.completed ? 'completed' : ''}`;
    taskItem.draggable = true; // Enable drag-and-drop
    taskItem.setAttribute('data-index', index); // Add an index attribute for reference

    taskItem.innerHTML = `
      <span>${task.name}</span>
      <div class="task-actions">
        <button class="complete">${task.completed ? 'Undo' : 'Complete'}</button>
        <button class="edit">Edit</button>
        <button class="delete">Delete</button>
      </div>
    `;

    // Mark task as complete
    taskItem.querySelector('.complete').addEventListener('click', () => {
      tasks[index].completed = !tasks[index].completed;
      saveAndRenderTasks();
    });

    // Edit task
    taskItem.querySelector('.edit').addEventListener('click', () => {
      const newTaskName = prompt('Edit task:', task.name);
      if (newTaskName) {
        tasks[index].name = newTaskName;
        saveAndRenderTasks();
      }
    });

    // Delete task
    taskItem.querySelector('.delete').addEventListener('click', () => {
      tasks.splice(index, 1);
      saveAndRenderTasks();
    });

    // Append task item to the list
    taskList.appendChild(taskItem);
  });

  // Update pending count
  pendingCount.textContent = tasks.filter(task => !task.completed).length;

  // Enable drag-and-drop functionality
  enableDragAndDrop();
}

// Function to save and render tasks
function saveAndRenderTasks() {
  localStorage.setItem('tasks', JSON.stringify(tasks));
  renderTasks();
}

// Drag-and-drop functionality
function enableDragAndDrop() {
  const taskItems = document.querySelectorAll('.task-item');

  taskItems.forEach(taskItem => {
    // Drag start event
    taskItem.addEventListener('dragstart', () => {
      taskItem.classList.add('dragging');
    });

    // Drag end event
    taskItem.addEventListener('dragend', () => {
      taskItem.classList.remove('dragging');
      reorderTasks(); // Update task order
    });
  });

  taskList.addEventListener('dragover', (e) => {
    e.preventDefault();
    const afterElement = getDragAfterElement(taskList, e.clientY);
    const draggingTask = document.querySelector('.dragging');

    if (afterElement == null) {
      taskList.appendChild(draggingTask);
    } else {
      taskList.insertBefore(draggingTask, afterElement);
    }
  });
}

// Helper function to get the element after which the dragged item should be placed
function getDragAfterElement(container, y) {
  const draggableElements = [...container.querySelectorAll('.task-item:not(.dragging)')];

  return draggableElements.reduce((closest, child) => {
    const box = child.getBoundingClientRect();
    const offset = y - box.top - box.height / 2;
    if (offset < 0 && offset > closest.offset) {
      return { offset, element: child };
    } else {
      return closest;
    }
  }, { offset: Number.NEGATIVE_INFINITY }).element;
}

// Function to reorder tasks based on the DOM order
function reorderTasks() {
  const reorderedTasks = [];
  document.querySelectorAll('.task-item').forEach(item => {
    const index = item.getAttribute('data-index');
    reorderedTasks.push(tasks[index]);
  });

  tasks = reorderedTasks;
  saveAndRenderTasks();
}
