const stars = document.getElementById("stars");
      for (let i = 0; i < 100; i++) {
        const star = document.createElement("div");
        star.className = "star";
        const xy = Math.random() * 100;
        const duration = Math.random() * 0.5 + 0.1;
        const size = Math.random() * 2;
        star.style.left = Math.random() * 100 + "vw";
        star.style.top = Math.random() * 100 + "vh";
        star.style.width = size * 50 + "px";
        star.style.height = size * 50 + "px";
        star.style.animationDuration = duration * 10 + "s";
        star.style.animationDelay = duration * -1 * Math.random() + "s";
        stars.appendChild(star);
        const colors = [
          "#ff0000",
          "#00ff00",
          "#0000ff",
          "#ffff00",
          "#ff00ff",
          "#00ffff",
        ];
        const randomColor = colors[Math.floor(Math.random() * colors.length)];
        star.style.backgroundColor = randomColor;
      }

      const form = document.querySelector("form");
      const message = document.getElementById("message");

      form.addEventListener("submit", async (e) => {
        e.preventDefault();
        const formData = new FormData(form);
        const response = await fetch("/api/register", {
          method: "POST",
          body: formData,
        });
        const data = await response.json();
        message.innerText = data.message;
      });