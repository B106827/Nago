import { createSelector } from 'reselect';

const messageSelector = (state) => state.message;

export const getMessageState = createSelector(
  [messageSelector],
  (state) => state.state
);

export const getMessageText = createSelector(
  [messageSelector],
  (state) => state.text
);

export const getMessageSeverity = createSelector(
  [messageSelector],
  (state) => state.severity
);
