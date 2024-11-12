import React from 'react';
import axios from 'axios';
import * as Yup from "yup";
import { Formik, Field, Form, ErrorMessage } from "formik";
import { FaTimes } from 'react-icons/fa';

const CreateUser = ({ createUser, setCreateUser, setUsers }) => {
  // Yup validation schema
  const validationSchema = Yup.object().shape({
    name: Yup.string()
      .min(3, "Name must be at least 3 characters")
      .required("Name is required"),
    email: Yup.string()
      .email("Invalid email address")
      .required("Email is required"),
    date_of_birth: Yup.date()
      .required("Birth date is required"),
  });

  const handleSubmit = (values, { resetForm }) => {
    const payload = { ...values };

    axios
      .post("http://localhost:8080/users", payload)
      .then((response) => {
        alert("User created successfully!");
        setUsers((prevUsers) => [...prevUsers, response.data]);
        resetForm();
        handleClose();
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const handleClose = () => {
    setCreateUser(false); // Close form
  };

  return (
    <>
      {createUser && (
        <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50 z-50">
          <div className="w-full max-w-2xl p-6 rounded-lg shadow-lg relative overflow-auto h-auto max-h-screen">
            <button className="absolute right-4 top-4" onClick={handleClose}>
              <FaTimes />
            </button>
            <h2 className="text-2xl font-bold mb-4">Create New User</h2>
            <Formik
              initialValues={{
                name: "",
                email: "",
                date_of_birth: "",
              }}
              validationSchema={validationSchema}
              onSubmit={handleSubmit}
            >
              {({ handleChange }) => (
                <Form noValidate>
                  <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
                    <div>
                      <label className="block mb-1">Name</label>
                      <Field
                        type="text"
                        name="name"
                        className="w-full px-3 py-2 border rounded-lg"
                        onChange={handleChange}
                      />
                      <ErrorMessage
                        name="name"
                        component="div"
                        className="text-red-500 text-sm"
                      />
                    </div>
                    <div>
                      <label className="block mb-1">Email</label>
                      <Field
                        type="email"
                        name="email"
                        className="w-full px-3 py-2 border rounded-lg"
                        onChange={handleChange}
                      />
                      <ErrorMessage
                        name="email"
                        component="div"
                        className="text-red-500 text-sm"
                      />
                    </div>
                    <div>
                      <label className="block mb-1">
                        Birth Date
                      </label>
                      <Field
                        type="date" // Updated to date input type
                        name="date_of_birth"
                        className="w-full px-3 py-2 border rounded-lg"
                        onChange={handleChange}
                      />
                      <ErrorMessage
                        name="date_of_birth"
                        component="div"
                        className="text-red-500 text-sm"
                      />
                    </div>
                  </div>

                  <button
                    type="submit"
                    className="bg-blue-500 text-white px-4 py-2 rounded-lg mt-4"
                  >
                    Create User
                  </button>
                </Form>
              )}
            </Formik>
          </div>
        </div>
      )}
    </>
  );
};

export default CreateUser;
