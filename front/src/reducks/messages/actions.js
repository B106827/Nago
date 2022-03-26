export const SHOW_MESSAGE = 'SHOW_MESSAGE';
export const showMessageAction = (severity = '', text = '') => {
  return {
    type: SHOW_MESSAGE,
    payload: {
      state: true,
      severity: severity,
      text: text,
    },
  };
};

export const HIDE_MESSAGE = 'HIDE_MESSAGE';
export const hideMessageAction = () => {
  return {
    type: HIDE_MESSAGE,
    payload: {
      state: false,
    },
  };
};
