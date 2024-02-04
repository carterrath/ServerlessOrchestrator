import { Link } from "react-router-dom";
import './../../assets/styles/App.css';

export function Home() {
  return (
    <div className="App">
    <div>
      <h1>Home</h1>
      <Link to="/Login">Login</Link>
    </div>
    </div>
  );
}