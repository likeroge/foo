import axios from "axios";
import type { ApiError } from "../entities/ApiError";

export const api = axios.create({
  // baseURL: "http://localhost:5000/api",
  baseURL: "/api",
});

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (!error.response) {
      const networkErr: ApiError = {
        message: "Сетевая ошибка. Проверьте подключение.",
      };
      return Promise.reject(networkErr);
    }
    // Response present
    const resp = error.response;
    const data = resp.data || {};
    // const traceId = resp.headers['x-request-id'] || data.trace_id;

    const apiError: ApiError = {
      message: data.message || resp.statusText,
      details: data.details,
    };

    // Доп. логика: 401 -> refresh token flow
    if (resp.status === 401) {
      // handle refresh or redirect to login
    }

    return Promise.reject(apiError);
  }
);
