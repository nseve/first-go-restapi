import { Routes, Route, Navigate } from "react-router-dom";

import { useState } from "react";

import Auth from "./pages/Auth";
import Projects from "./pages/Projects";
import Tasks from "./pages/Tasks";

export default function App() {
  const [token, setToken] = useState(
    localStorage.getItem("token")
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
    return (
      <Routes>
        <Route
          path="/auth"
          element={<Auth onLogin={handleLogin} />}
        />

        <Route
          path="*"
          element={<Navigate to="/auth" />}
        />
      </Routes>
    );
  }

  return (
    <Routes>
      <Route
        path="/projects"
        element={
          <Projects
            token={token}
            onLogout={handleLogout}
          />
        }
      />

      <Route
        path="/projects/:id/tasks"
        element={<Tasks token={token} />}
      />

      <Route
        path="*"
        element={<Navigate to="/projects" />}
      />
    </Routes>
  );
}