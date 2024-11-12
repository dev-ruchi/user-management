import React, { useState, useEffect } from "react";
import axios from "axios";
import { Link } from "react-router-dom";
import { Edit, Trash } from "react-feather";
import CreateUser from "./CreateUser";

const UserList = () => {
  const [users, setUsers] = useState([]);
  const [updateUserIndex, setupdateUserIndex] = useState(-1);
  const [deleteUserId, setDeleteUserId] = useState(null);
  const [createUser, setCreateUser] = useState(false);

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = () => {
    axios.get(`http://localhost:8080/users`).then((response) => {
      setUsers(response.data);
      console.log(response.data);
    });
  };

  const handleCreateClick = () => {
    setCreateUser(true); // Open form when "Create" button is clicked
  };

  return (
    <div className="p-6">
      <div className="flex justify-between items-center mb-4">
        <h2 className="text-3xl font-bold text-primary">Users</h2>
        <button
          onClick={handleCreateClick}
          className="btn btn-primary px-6 py-2"
        >
          Create
        </button>
      </div>
      <div className="overflow-x-auto">
        <table className="table w-full">
          <thead>
            <tr className="bg-neutral text-neutral-content">
              <th className="py-3 px-4 text-left sticky left-0 z-10">Name</th>
              <th className="py-3 px-4 text-left">Email</th>
              <th className="py-3 px-4 text-left">Birth Date</th>
              <th className="py-3 px-4 text-left">Actions</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user, index) => (
              <tr
                key={`${user.id}-${index}`}
                className="hover:bg-gray-100 even:bg-base-200 odd:bg-base-100"
              >
                <td className="py-3 px-4 text-left sticky left-0 z-10">
                  <Link
                    to={`/users/${user.id}`}
                    className="text-blue-500 hover:underline"
                  >
                    {user.name}
                  </Link>
                </td>
                <td className="py-3 px-4 text-left">{user.email}</td>
                <td className="py-3 px-4 text-left">{user.date_of_birth}</td>
                <td className="py-3 px-4 text-left">
                  <button
                    onClick={() => setupdateUserIndex(index)}
                    title="Edit"
                  >
                    <Edit className="mr-2" />
                  </button>
                  <button
                    onClick={() => setDeleteUserId(user.id)}
                    title="Delete"
                  >
                    <Trash className="mr-2" />
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <CreateUser
        createUser={createUser}
        setCreateUser={setCreateUser}
        setUsers={setUsers}
      />
    </div>
  );
};

export default UserList;
