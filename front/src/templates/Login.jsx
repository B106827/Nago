import { useCallback, useState, useEffect } from 'react';
import { TextInput, PrimaryButton } from '../components/UIkit';
import { login } from '../reducks/users/operations';
import { useDispatch, useSelector } from 'react-redux';
import { push } from 'connected-react-router';
import { getIsLogedIn } from '../reducks/users/selectors';

const Login = () => {
  const dispatch = useDispatch();
  const selector = useSelector((state) => state);
  const isLogedIn = getIsLogedIn(selector);
  useEffect(() => {
    if (isLogedIn) {
      dispatch(push('/'));
    }
  });

  const [email, setEmail]   = useState(''),
    [password, setPassword] = useState('');

  const inputEmail = useCallback(
    (event) => {
      setEmail(event.target.value);
    },
    [setEmail]
  );

  const inputPassword = useCallback(
    (event) => {
      setPassword(event.target.value);
    },
    [setPassword]
  );

  return (
    <div className='c-section-container'>
      <h2 className='u-text__headline__k u-text-center'>ログイン</h2>
      <div className='module-spacer--short' />
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
      <TextInput
        fullWidth={true}
        label={'パスワード'}
        multiline={false}
        required={true}
        rows={1}
        value={password}
        type={'password'}
        onChange={inputPassword}
      />
      <div className='module-spacer--medium' />
      <div className='center'>
        <PrimaryButton
          label={'ログイン'}
          onClick={() => dispatch(login(email, password))}
          addStyle={{ fontWeight: 'bold' }}
        />
        <div className='module-spacer--medium' />
        <p
          className='cursor-pointer'
          onClick={() => dispatch(push('/email_register'))}
        >
          新規登録はこちら
        </p>
        <p
          className='cursor-pointer'
          onClick={() => dispatch(push('/password_reset'))}
        >
          パスワードを忘れた方はこちら
        </p>
      </div>
    </div>
  );
};

export default Login;
