import { useState } from "react";
import Auth from "./pages/Auth";
import Projects from "./pages/Projects";

export default function App() {
  const [token, setToken] = useState(
    localStorage.getItem("token") || null
  );

  const handleLogin = (token) => {
    localStorage.setItem("token", token);
    setToken(token);
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    setToken(null);
  };

  if (!token) {
    return <Auth onLogin={handleLogin} />;
  }

  return <Projects token={token} onLogout={handleLogout} />;
}