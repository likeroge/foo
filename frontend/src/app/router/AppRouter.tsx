import { createBrowserRouter } from "react-router";
import { AboutPage, HomePage } from "../../pages";
import { AddUserPage, AllUsersPage, FindUserPage } from "../../pages/users";

export const appRouter = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/about",
    element: <AboutPage />,
  },
  {
    path: "/users/list",
    element: <AllUsersPage />,
  },
  {
    path: "/users/add",
    element: <AddUserPage />,
  },
  {
    path: "/users/find",
    element: <FindUserPage />,
  },
]);
