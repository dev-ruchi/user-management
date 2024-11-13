import React from "react";
import { FaTimes } from "react-icons/fa";
import { backend } from "../../backend";

const DeleteUser = ({ userId, onDelete, onClose }) => {
  const handleDelete = () => {
    backend
      .delete(`http://localhost:8080/users/${userId}`)
      .then((response) => {
        if (!response.ok) {
          throw new Error("Could not delete the user");
        }
        onDelete(userId); // Notify parent to update the list
      })
      .catch((err) => {
        console.error("Error deleting user:", err);
      });
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 z-50">
      <div className="w-full max-w-lg p-6 rounded-lg shadow-lg relative overflow-auto bg-base-100">
        <button
          className="absolute right-4 top-4 text-gray-500 hover:text-gray-700"
          onClick={onClose}
        >
          <FaTimes />
        </button>
        <h2 className="text-3xl font-bold mb-4 text-center text-primary">
          Delete User
        </h2>
        <p className="text-center mb-4">
          Are you sure you want to delete this user?
        </p>
        <div className="flex justify-center gap-4">
          <button
            onClick={handleDelete}
            className="btn btn-danger text-white px-4 py-2 rounded"
          >
            Delete
          </button>
          <button
            onClick={onClose}
            className="btn btn-secondary text-white px-4 py-2 rounded"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  );
};

export default DeleteUser;
