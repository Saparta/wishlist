const API = import.meta.env.VITE_API_BASE;

export async function me() {
  const res = await fetch(`${API}/auth/me`, { credentials: 'include' });
  return res.json(); // { user: ... | null }
}

export function startGoogleLogin() {
  // Let the server build the URL and redirect us
  window.location.href = `${API}/auth/google/start`;
}

export async function logout() {
  await fetch(`${API}/auth/logout`, { method: 'POST', credentials: 'include' });
}