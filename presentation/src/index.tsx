import { Router } from './router';
//import './assets/styles/main.less'
import './assets/styles/styles.css';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="min-h-svh">
        <Router />
      </div>
    </QueryClientProvider>
  );
}

export default App;
