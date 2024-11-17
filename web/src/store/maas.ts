import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import { IMaas } from 'src/types';

type MaasState = {
  maasList: IMaas[];
};

type MaasAction = {
  setMaasList: (maasList: IMaas[]) => void;
};

const initialState: MaasState = {
  maasList: [],
};

export const useMaasStore = create<MaasState & MaasAction>()(
  devtools((set, get) => ({
    ...initialState,
    setMaasList(maasList) {
      set({ maasList });
    },
  }))
);
