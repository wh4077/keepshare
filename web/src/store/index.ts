import { create } from "zustand";
import { AccountState, createAccountStore } from "./account";
import { GlobalState, createGlobalStore } from "./global";
import { LinkState, createLinkStore } from "./link";
import { PikPakState, createPikPakStore } from "./pikpak";
import { RapidGatorState, createRapidGatorStore } from "./rapidgator";

const useStore = create<GlobalState & AccountState & LinkState & PikPakState & RapidGatorState>((...set) => ({
  ...createAccountStore(...set),
  ...createGlobalStore(...set),
  ...createLinkStore(...set),
  ...createPikPakStore(...set),
  ...createRapidGatorStore(...set),
}));

export default useStore;
