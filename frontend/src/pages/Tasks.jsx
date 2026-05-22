import { useEffect, useState } from "react";

import { useNavigate, useParams } from "react-router-dom";

import {
  getTasks,
  createTask,
  updateTask,
  deleteTask,
} from "../api";

export default function Tasks({ token }) {
  const navigate = useNavigate();

  const { id } = useParams();

  const projectId = Number(id);

  const [tasks, setTasks] = useState([]);

  const [title, setTitle] = useState("");
  const [duration, setDuration] = useState("");

  const [editingId, setEditingId] = useState(null);
  const [editingTitle, setEditingTitle] = useState("");
  const [editingDuration, setEditingDuration] = useState("");

  const [selectedTaskId, setSelectedTaskId] = useState(null);

  const [searchId, setSearchId] = useState("");

  const [error, setError] = useState("");

  const load = async () => {
    try {
      const data = await getTasks(token, projectId);

      setError("");
      setTasks(data);
    } catch (err) {
      setError(err.message);
    }
  };

  useEffect(() => {
    load();
  }, []);

  const handleCreate = async () => {
    try {
      await createTask(
        token,
        projectId,
        title,
        Number(duration)
      );

      setTitle("");
      setDuration("");

      load();
    } catch (err) {
      setError(err.message);
    }
  };

  const handleUpdate = async (id) => {
    try {
      await updateTask(
        token,
        id,
        editingTitle,
        Number(editingDuration)
      );

      setEditingId(null);

      load();
    } catch (err) {
      setError(err.message);
    }
  };

  const handleDelete = async () => {
    if (!selectedTaskId) return;

    try {
      await deleteTask(token, selectedTaskId);

      setSelectedTaskId(null);

      load();
    } catch (err) {
      setError(err.message);
    }
  };

  const handleSearch = async () => {
    if (!searchId) {
      load();
      return;
    }

    try {
      const res = await fetch(
        `http://localhost:8080/tasks/${searchId}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error);
      }

      setError("");
      setTasks([data]);
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="container">
      <button
        className="back"
        onClick={() => navigate("/projects")}
      >
        ← Back
      </button>

      <h2>Project Tasks</h2>

      <div className="form">
        <input
          placeholder="Task title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />

        <input
          type="number"
          placeholder="Duration"
          value={duration}
          onChange={(e) => setDuration(e.target.value)}
        />

        <button onClick={handleCreate}>
          Add Task
        </button>
      </div>

      <div className="search-block">
        <input
          placeholder="Search task by ID"
          value={searchId}
          onChange={(e) => setSearchId(e.target.value)}
        />

        <button onClick={handleSearch}>
          Search
        </button>
      </div>

      {error && <p className="error">{error}</p>}

      <ul>
        {tasks.map((t) => (
          <li
            key={t.id}
            className={
              selectedTaskId === t.id
                ? "selected"
                : ""
            }
            onClick={() => setSelectedTaskId(t.id)}
          >
            <div className="row">
              {editingId === t.id ? (
                <>
                  <input
                    value={editingTitle}
                    onChange={(e) =>
                      setEditingTitle(e.target.value)
                    }
                  />

                  <input
                    type="number"
                    value={editingDuration}
                    onChange={(e) =>
                      setEditingDuration(e.target.value)
                    }
                  />

                  <button
                    className="small"
                    onClick={() =>
                      handleUpdate(t.id)
                    }
                  >
                    Save
                  </button>
                </>
              ) : (
                <>
                  <div className="item-content">
                    <div className="item-title">
                      {t.title}
                    </div>

                    <div className="item-meta">
                      ID: {t.id} • [{t.duration} min]
                    </div>
                  </div>

                  <button
                    className="small"
                    onClick={() => {
                      setEditingId(t.id);
                      setEditingTitle(t.title);
                      setEditingDuration(t.duration);
                    }}
                  >
                    Edit
                  </button>
                </>
              )}
            </div>
          </li>
        ))}
      </ul>

      <button
        className="danger"
        onClick={handleDelete}
      >
        Delete
      </button>
    </div>
  );
}