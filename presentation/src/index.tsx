import { Router } from './router'
import './assets/styles/main.less'
import { NavBar } from './components/NavBar';

function App() {
  
  return (
    <div className="main-container">
      <NavBar />
      <div className="main-body" id="body">
        <Router />
      </div>
    </div>
    );
  }
  
  export default App;