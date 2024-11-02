import { get } from 'lodash';

import axios, { endpoints, setSession } from 'src/utils/axios';

import { CONFIG } from 'src/config-global';

// ----------------------------------------------------------------------

export type SignInParams = {
  username: string;
  password: string;
};

export type SignUpParams = {
  email: string;
  password: string;
  username: string;
};

/** **************************************
 * Sign in
 *************************************** */
export const signInWithPassword = async ({ username, password }: SignInParams): Promise<void> => {
  try {
    const params = { username, password };

    const res = await axios.post(endpoints.auth.signIn, params);

    // const { accessToken } = res.data;
    // const { token: accessToken } = res?.data?.header;

    const accessToken = get(res, 'data.header.token', '');

    if (!accessToken) {
      throw new Error('Access token not found in response');
    }

    await setSession(accessToken);
  } catch (error) {
    console.error('Error during sign in:', error);
    throw error;
  }
};

/** **************************************
 * Sign up
 *************************************** */
export const signUp = async ({ email, password, username }: SignUpParams): Promise<void> => {
  const params = {
    email,
    password,
    username,
  };

  try {
    const res = await axios.post(endpoints.auth.signUp, params);

    const { token: accessToken } = res.data.header;

    if (!accessToken) {
      throw new Error('Access token not found in response');
    }

    sessionStorage.setItem(CONFIG.ACCESS_TOKEN, accessToken);
  } catch (error) {
    console.error('Error during sign up:', error);
    throw error;
  }
};

/** **************************************
 * Sign out
 *************************************** */
export const signOut = async (): Promise<void> => {
  try {
    await setSession(null);
  } catch (error) {
    console.error('Error during sign out:', error);
    throw error;
  }
};
