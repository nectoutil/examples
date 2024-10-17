import React from 'react';

const App: React.FC = () => {
  return (
    <div className="app">
      <h1>Hello from Vite, SWC, Styled JSX, and TypeScript!</h1>
      <style jsx>{`
        .app {
          display: flex;
          justify-content: center;
          align-items: center;
          height: 100vh;
          font-family: Arial, sans-serif;
          background-color: #f0f0f0;
        }
      `}</style>
    </div>
  );
};

export default App;
