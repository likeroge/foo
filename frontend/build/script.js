console.log("Hello world!");

async function changeColor() {
  console.log("changeColor func3");

  let data = await fetch("/api/users");

  console.log(data);
  const users = await data.json();
  console.log(users);

  const title = document.querySelector(".main-title");
  const colors = ["#ff6b6b", "#4ecdc4", "#45b7d1", "#f9ca24", "#6c5ce7"];
  const randomColor = colors[Math.floor(Math.random() * colors.length)];
  title.style.color = randomColor;
  title.style.textShadow = `3px 3px 6px ${randomColor}40`;
}

function addSparkles() {
  const title = document.querySelector(".main-title");
  const rect = title.getBoundingClientRect();

  for (let i = 0; i < 20; i++) {
    createSparkle(
      rect.left + Math.random() * rect.width,
      rect.top + Math.random() * rect.height
    );
  }
}

function createSparkle(x, y) {
  const sparkle = document.createElement("div");
  sparkle.className = "sparkle";
  sparkle.style.left = `${x}px`;
  sparkle.style.top = `${y}px`;
  document.body.appendChild(sparkle);

  setTimeout(() => {
    sparkle.remove();
  }, 1000);
}

// Добавляем анимацию при загрузке страницы
document.addEventListener("DOMContentLoaded", function () {
  const title = document.querySelector(".main-title");
  title.style.opacity = "0";
  title.style.transform = "translateY(-30px)";
  title.style.transform = "rotate(359deg)";

  setTimeout(() => {
    title.style.transition = "all 0.8s ease";
    title.style.opacity = "1";
    title.style.transform = "translateY(0)";
  }, 100);
});
