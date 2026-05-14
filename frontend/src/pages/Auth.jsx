import { useState } from "react";
import { login, register } from "../api";

export default function Auth({ onLogin }) {
  const [isLogin, setIsLogin] = useState(true);

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const [error, setError] = useState("");

  const handleSubmit = async () => {
    try {
      setError("");

      let data;

      if (isLogin) {
        data = await login(email, password);
      } else {
        data = await register(email, password);
      }

      if (data.token) {
        onLogin(data.token);
      } else {
        setIsLogin(true);
      }
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="container">
      <h2>{isLogin ? "Login" : "Register"}</h2>

      <input
        placeholder="Email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
      />

      <input
        type="password"
        placeholder="Password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />

      <button onClick={handleSubmit}>
        {isLogin ? "Login" : "Register"}
      </button>

      <p className="switch-text">
        {isLogin ? "No account?" : "Already have an account?"}
      </p>

      <button
        className="switch-button"
        onClick={() => {
          setError("");
          setIsLogin(!isLogin);
        }}
      >
        {isLogin ? "Go to Register" : "Go to Login"}
      </button>

      {error && <p className="error">{error}</p>}
    </div>
  );
}