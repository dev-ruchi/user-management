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
      <h2 className="text-3xl">Users</h2>
      <button
          onClick={handleCreateClick}
          className="bg-blue-500 text-white px-6 py-4 rounded"
        >
          Create
        </button>
      <table>
        <thead>
          <tr className="uppercase text-sm leading-normal">
            <th className="py-3 px-3 text-left sticky left-0 z-10">
              Name
            </th>
            <th className="py-3 px-3 text-left">Email</th>
            <th className="py-3 px-3 text-left">BirthDate</th>
            <th className="py-3 px-3 text-left">Actions</th>
          </tr>
        </thead>
        <tbody className="text-gray-600 text-sm font-light">
          {users.map((user, index) => (
            <tr
              key={`${user.id}-${index}`}
              className={`border-b border-gray-200 hover:bg-gray-100`}
            >
              <td className="py-3 px-3 text-left whitespace-nowrap sticky left-0 z-10">
                <Link
                  to={`/users/${user.id}`}
                  className="text-blue-500 hover:underline"
                >
                  {user.name}
                </Link>
              </td>
              <td className="py-3 px-3 text-left">{user.email}</td>
              <td className="py-3 px-3 text-left">{user.date_of_birth}</td>
              
            </tr>
          ))}
        </tbody>
      </table>
      <CreateUser
        createUser={createUser}
        setCreateUser={setCreateUser}
        setUsers={setUsers}
      />
    </div>

    
  );
};

export default UserList;
