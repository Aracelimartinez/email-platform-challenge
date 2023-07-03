/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html",
  "./src/**/*.{vue,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        "color-primary": "#d52a47",
        "color-secondary": "#df5134",
      }
    },
    fontFamily: {
      Roboto: ["Roboto", "sans-serif"],
    },
    container: {
      padding: "2rem",
      center: true,
    }
  },
  plugins: [],
}
