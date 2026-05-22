import { useEffect, useState } from "react";

import { useNavigate } from "react-router-dom";

import {
  getProjects,
  createProject,
  updateProject,
  deleteProject,
} from "../api";

export default function Projects({ token, onLogout }) {
  const navigate = useNavigate();

  const [projects, setProjects] = useState([]);

  const [title, setTitle] = useState("");

  const [selectedProjectId, setSelectedProjectId] =
    useState(null);

  const [editingId, setEditingId] = useState(null);
  const [editingTitle, setEditingTitle] =
    useState("");

  const [searchId, setSearchId] = useState("");

  const [error, setError] = useState("");

  const load = async () => {
    try {
      const data = await getProjects(token);

      setError("");
      setProjects(data);
    } catch (err) {
      setError(err.message);
    }
  };

  useEffect(() => {
    load();
  }, []);

  const handleCreate = async () => {
    try {
      await createProject(token, title);

      setTitle("");

      load();
    } catch (err) {
      setError(err.message);
    }
  };

  const handleUpdate = async (id) => {
    try {
      await updateProject(token, id, editingTitle);

      setEditingId(null);
      setEditingTitle("");

      load();
    } catch (err) {
      setError(err.message);
    }
  };

  const handleDelete = async () => {
    if (!selectedProjectId) return;

    try {
      await deleteProject(token, selectedProjectId);

      setSelectedProjectId(null);

      load();
    } catch (err) {
      setError(err.message);
    }
  };

  const handleSearch = async () => {
    if (!searchId) {
      setError("");
      load();
      return;
    }

    try {
      const res = await fetch(
        `http://localhost:8080/projects/${searchId}`,
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
      setProjects([data]);
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="container">
      <div className="header">
        <h2>User Projects</h2>

        <button
          className="logout"
          onClick={onLogout}
        >
          Logout
        </button>
      </div>

      <div className="form">
        <input
          value={title}
          onChange={(e) =>
            setTitle(e.target.value)
          }
          placeholder="New project"
        />

        <button onClick={handleCreate}>
          Create
        </button>
      </div>

      <div className="search-block">
        <input
          placeholder="Search project by ID"
          value={searchId}
          onChange={(e) =>
            setSearchId(e.target.value)
          }
        />

        <button onClick={handleSearch}>
          Search
        </button>
      </div>

      {error && <p className="error">{error}</p>}

      <ul>
        {projects.map((p) => (
          <li
            key={p.id}
            className={
              selectedProjectId === p.id
                ? "selected"
                : ""
            }
            onClick={() => setSelectedProjectId(p.id)}
          >
            <div className="row">
              {editingId === p.id ? (
                <>
                  <input
                    value={editingTitle}
                    onChange={(e) =>
                      setEditingTitle(
                        e.target.value
                      )
                    }
                  />

                  <button
                    className="small"
                    onClick={() =>
                      handleUpdate(p.id)
                    }
                  >
                    Save
                  </button>
                </>
              ) : (
                <>
                  <div
                    className="project-info"
                  >
                    <div className="item-content">
                      <div className="item-title">
                        {p.title}
                      </div>

                      <div className="item-meta">
                        ID: {p.id}
                      </div>
                    </div>
                  </div>

                  <div className="actions">
                    <button
                      className="small"
                      onClick={(e) => {
                        e.stopPropagation();
                        navigate(`/projects/${p.id}/tasks`)
                      }}
                    >
                      Open
                    </button>

                    <button
                      className="small"
                      onClick={() => {
                        setEditingId(p.id);
                        setEditingTitle(
                          p.title
                        );
                      }}
                    >
                      Edit
                    </button>
                  </div>
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