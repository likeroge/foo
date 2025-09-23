import { createBrowserRouter } from "react-router";
import { AboutPage, HomePage } from "../../pages";

export const appRouter = createBrowserRouter([
  {
    path: "/",
    element: <HomePage />,
  },
  {
    path: "/about",
    element: <AboutPage />,
  },
]);
