import { axiosWrapper } from ".";

/* RapidGator account management */

// get pikpak host info
export interface RapidGatorHostInfo {
  // master: {
  //   user_id: string;
  //   keepshare_user_id: string;
  //   email: string;
  //   password: string;
  //   created_at: string;
  //   updated_at: string;
  // };
  // revenue: number;
  // workers: {
  //   premium: {
  //     count: number;
  //     used: number;
  //     limit: number;
  //   };
  //   free: {
  //     count: number;
  //     used: number;
  //     limit: number;
  //   };
  // };
  account: string;
}
export const getRapidGatorHostInfo = () => {
  return axiosWrapper<RapidGatorHostInfo>({
    url: "/api/host/info?host=rapidgator",
    method: "GET",
  });
};

export interface SetAccountParams {
  account: string;
  password: string;
  // if action == "set", RapidGator will set account and password;
  action:string;
}
export interface SetAccountResponse {
  success: boolean;
}
export const setRapidGatorHostInfo = (params: SetAccountParams)  => {
  return axiosWrapper<SetAccountResponse>({
    url: "/api/host/info?host=rapidgator",
    method: "POST",
    data: params,
  });
};

export interface GetRapidGatorAccountStatisticsParams {
  host: string;
  stored_count_lt: number[];
  not_stored_days_gt: number[];
}
export type GetRapidGatorAccountStatisticsResponse = Record<
  "stored_count_lt" | "not_stored_days_gt",
  Array<{
    number: number;
    total_count: number;
    total_size: number;
  }>
>;
// get account storage statistics
export const getRapidGatorAccountStatistics = (
  params: GetRapidGatorAccountStatisticsParams,
) => {
  return axiosWrapper<GetRapidGatorAccountStatisticsResponse>({
    url: "/api/storage/statistics",
    method: "POST",
    data: params,
  });
};

export interface ClearRapidGatorAccountStorageParams {
  host: string;
  stored_count_lt?: number;
  not_stored_days_gt?: number;
  only_for_premium?: boolean;
}
// clear account usage storage
export const clearRapidGatorAccountStorage = (
  params: ClearRapidGatorAccountStorageParams,
) => {
  return axiosWrapper({
    url: "/api/storage/release",
    method: "POST",
    data: params,
  });
};
