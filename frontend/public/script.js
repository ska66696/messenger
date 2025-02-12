document.addEventListener('DOMContentLoaded', function () {
    const registrationForm = document.getElementById('registrationForm');
    const registrationResultDiv = document.getElementById('registrationResult');

    registrationForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const username = document.getElementById('regUsername').value;
        const email = document.getElementById('regEmail').value;
        const password = document.getElementById('regPassword').value;

        fetch('http://localhost:8080/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                email: email,
                password: password
            })
        })
            .then(response => response.json())
            .then(data => {
                if (data.status === 'success') {
                    registrationResultDiv.textContent = 'Регистрация успешна!';
                    registrationResultDiv.style.color = 'green';
                } else {
                    registrationResultDiv.textContent = 'Ошибка регистрации: ' + data.message;
                    registrationResultDiv.style.color = 'red';
                }
            })
            .catch(error => {
                registrationResultDiv.textContent = 'Ошибка сети: ' + error;
                registrationResultDiv.style.color = 'red';
            });
    });

    const loginForm = document.getElementById('loginForm');
    const loginResultDiv = document.getElementById('loginResult');

    loginForm.addEventListener('submit', function (event) {
        event.preventDefault();

        const username = document.getElementById('loginUsername').value;
        const password = document.getElementById('loginPassword').value;

        fetch('http://localhost:8080/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        })
            .then(response => response.json())
            .then(data => {
                if (data.status === 'success') {
                    loginResultDiv.textContent = 'Вход выполнен! User ID: ' + data.user_id + ', Username: ' + data.username;
                    loginResultDiv.style.color = 'green';
                } else {
                    loginResultDiv.textContent = 'Ошибка входа: ' + data.message;
                    loginResultDiv.style.color = 'red';
                }
            })
            .catch(error => {
                loginResultDiv.textContent = 'Ошибка сети: ' + error;
                loginResultDiv.style.color = 'red';
            });
    });
});