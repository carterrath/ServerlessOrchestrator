/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      animation: {
        bounce200: 'bounce 1s infinite 200ms',
        'slide-up': 'slide-up 0.3s ease-out forwards',
      },
      keyframes: {
        bounce: {
          '0%, 100%': { transform: 'translateY(-25%)', animationTimingFunction: 'cubic-bezier(0.8,0,1,1)' },
          '50%': { transform: 'none', animationTimingFunction: 'cubic-bezier(0,0,0.2,1)' },
        },
        'slide-up': {
          '0%': { transform: 'translateY(100%)' },
          '100%': { transform: 'translateY(0)' },
        },
      },
      colors: {
        darkPink: '#d04466',
        amaranthPink: '#f194b4',
        caribbeanCurrent: '#006c67',
        lightTeal: '#41bba6',
        midnightGreen: '#003844',
        selectiveYellow: '#ffb100',
        dutchWhite: '#ffebc6',
      },
    },
  },
  plugins: [],
};
