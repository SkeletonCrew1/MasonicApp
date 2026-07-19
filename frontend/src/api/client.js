import axios from "axios";

const host = window.location.hostname;

export const django = axios.create({
  baseURL: `http://${host}:8000/api`,
});

export const banService = axios.create({
  baseURL: `http://${host}:8081`,
});

export const broadcastService = axios.create({
  baseURL: `http://${host}:8082`,
});
export const fetchPosts = () => django.get("/posts/").then((r) => r.data);
export const createPost = (payload) =>
  django.post("/posts/", payload).then((r) => r.data);

export const searchUsers = (q) =>
  django.get("/users/search/", { params: { q } }).then((r) => r.data);
export const promoteUser = (userId, status) =>
  django.post(`/users/${userId}/promote/`, { status }).then((r) => r.data);
export const inviteUser = (email) =>
  django.post("/invite/", { email }).then((r) => r.data);
export const deleteAllData = () => django.post("/admin/delete-all/");
export const banIp = (ip) => banService.post("/ban", { ip }).then((r) => r.data);
export const checkIp = (ip) =>
  banService.get("/check", { params: { ip } }).then((r) => r.data);
export const listBans = () => banService.get("/bans").then((r) => r.data);

export const sendBroadcast = (message, statuses) =>
  broadcastService
    .post("/broadcast", { message, statuses })
    .then((r) => r.data);
export const listBroadcasts = () =>
  broadcastService.get("/broadcasts").then((r) => r.data);
