import { RapidGatorHostInfo } from "@/api/rapidgator";
import { StateCreator } from "zustand";

export interface RapidGatorState {
	rapidGatorInfo: Partial<RapidGatorHostInfo>;

	setRapidGatorInfo: (rapidGatorInfo: RapidGatorHostInfo) => void;
}

export const createRapidGatorStore: StateCreator<RapidGatorState> = (set) => ({
	rapidGatorInfo: {},

	setRapidGatorInfo: (rapidGatorInfo) => set({ rapidGatorInfo }),
});
