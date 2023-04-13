/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
    "./app/**/*.{js,ts,jsx,tsx}"
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))"
      },
      colors: {
        mmsPry3: "#058B94",
        mmsBlack1: "#141414",
        mmsBlack2: "#333333",
        mmsBlack3: "#4D4D4D",
        mmsPry10: "#E6FDFE",
        mmsBlack5: "#808080",
        green11: "#F7FEFF"
      }
    }
  },
  plugins: [require("daisyui")]
};
