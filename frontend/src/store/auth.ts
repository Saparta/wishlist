import { create } from 'zustand';
import { me } from '../lib/api';

type User = { id: string; email: string; name?: string; avatar_url?: string } | null;

export const useAuth = create<{
  user: User;
  loading: boolean;
  load: () => Promise<void>;
  setUser: (u: User) => void;
}>((set) => ({
  user: null,
  loading: true,
  setUser: (u) => set({ user: u }),
  load: async () => {
    try {
      const data = await me();
      set({ user: data.user ?? null });
    } finally {
      set({ loading: false });
    }
  },
}));