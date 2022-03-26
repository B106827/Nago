import { useCallback, useState, useEffect } from 'react';
import { TextInput, PrimaryButton } from '../components/UIkit';
import { resetPassword } from '../reducks/users/operations';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';
import { getIsLogedIn } from '../reducks/users/selectors';

const PasswordReset = () => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);
  useEffect(() => {
    if (isLogedIn) {
      dispatch(push('/'));
    }
  });

  const [email, setEmail] = useState('');

  const inputEmail = useCallback(
    (event) => {
      setEmail(event.target.value);
    },
    [setEmail]
  );

  return (
    <div className='c-section-container'>
      <h2 className='u-text__headline__k u-text-center'>パスワードリセット</h2>
      <div className='module-spacer--medium' />
      <TextInput
        fullWidth={true}
        label={'メールアドレス'}
        multiline={false}
        required={true}
        rows={1}
        value={email}
        type={'email'}
        onChange={inputEmail}
      />
      <div className='module-spacer--medium' />
      <div className='center'>
        <PrimaryButton
          label={'パスワードリセット'}
          onClick={() => dispatch(resetPassword(email))}
          addStyle={{ fontWeight: 'bold' }}
        />
        <div className='module-spacer--medium' />
        <p className='cursor-pointer' onClick={() => dispatch(push('/login'))}>
          ログイン画面に戻る
        </p>
      </div>
    </div>
  );
};

export default PasswordReset;
