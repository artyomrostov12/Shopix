
document.addEventListener("DOMContentLoaded", function () {
    // Находим все кнопки с атрибутом data-action
    const buttons = document.querySelectorAll("[data-action]");

    // Перебираем кнопки
    buttons.forEach(function (button) {
        button.addEventListener("click", function () {
            const action = button.dataset.action;

            switch (action) {
                case "home":
                    window.location.href = "../index.html";
                    break;

                case "cart":
                    alert("Корзина пока находится в разработке");
                    break;

                case "profile":
                    window.location.href = "profile.html";
                    break;

                case "orders":
                    alert("Здесь будут отображаться ваши заказы");
                    break;

                case "favorites":
                    alert("Здесь будут отображаться избранные товары");
                    break;

                case "settings":
                    alert("Раздел настроек пока в разработке");
                    break;

                case "logout":
                    const confirmLogout = confirm(
                        "Вы действительно хотите выйти из аккаунта?"
                    );

                    if (confirmLogout) {
                        alert("Вы вышли из аккаунта");
                    }
                    break;
            }
        });
    });
});

document.addEventListener("DOMContentLoaded", function () {
    // Находим все кнопки с атрибутом data-action
    const buttons = document.querySelectorAll("[data-action]");

    // Перебираем кнопки
    buttons.forEach(function (button) {
        button.addEventListener("click", function () {
            const action = button.dataset.action;

            switch (action) {
                case "home":
                    window.location.href = "../index.html";
                    break;

                case "cart":
                    alert("Корзина пока находится в разработке");
                    break;

                case "profile":
                    window.location.href = "profile.html";
                    break;

                case "orders":
                    alert("Здесь будут отображаться ваши заказы");
                    break;

                case "favorites":
                    alert("Здесь будут отображаться избранные товары");
                    break;

                case "settings":
                    alert("Раздел настроек пока в разработке");
                    break;

                case "logout":
                    const confirmLogout = confirm(
                        "Вы действительно хотите выйти из аккаунта?"
                    );

                    if (confirmLogout) {
                        alert("Вы вышли из аккаунта");
                    }
                    break;
            }
        });
    });
});
