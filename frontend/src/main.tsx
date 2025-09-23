import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { AppLayout } from "./app/layouts/AppLayout.tsx";
import { RouterProvider } from "react-router";
import { appRouter } from "./app/router/AppRouter.tsx";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AppLayout>
      <RouterProvider router={appRouter} />
    </AppLayout>
  </StrictMode>
);
