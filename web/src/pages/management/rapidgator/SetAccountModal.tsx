import { useState } from 'react';
import { Modal, Form, message } from "antd";
import { setRapidGatorHostInfo } from '@/api/rapidgator';

import { StyledForm, StyledInput, PasswordInput } from './style';

interface FieldType {
  account?: string;
  password?: string;
}

interface SetAccountModalProps {
  open: boolean;
  onClose: () => void;
}

function SetAccountModal(props: SetAccountModalProps) {
  const { open } = props;
  const [setAccountConfirm, setSetAccountConfirm] = useState(false);
  const [isSetAccountLoading, setIsSetAccountLoading] = useState(false);
  const [form] = Form.useForm<FieldType>();

  const handleSetAccountModalClose = () => {
    props.onClose && props.onClose();
    form.resetFields();
  };

  const handleSetAccountFormChange = async () => {
    try {
      await form.validateFields({
        validateOnly: true,
      })
      setSetAccountConfirm(true);
    } catch {
      setSetAccountConfirm(false);
    }
  };

  const handleSetAccount = async ({ account, password }: FieldType) => {
    if (!account || !password) {
      console.error('can not find form field.');
      message.error('can not find form field.');
      return;
    }

    setIsSetAccountLoading(true);
    const result = await setRapidGatorHostInfo({
      account: account!,
      password: password!,
      action: "set",
    });
    setIsSetAccountLoading(false);

    if (result.error) {
      message.error(result.error.message || 'set account fail');
      return;
    }

    form.resetFields();

    // an ugly implementation, to refresh account info in page(maybe should use useStore?), and close modal;
    window.location.reload();
  };

  return (
    <Modal
      title="Set RapidGator Account"
      okText="Confirm"
      open={open}
      onCancel={handleSetAccountModalClose}
      onOk={() => form.submit()}
      okButtonProps={{
        disabled: !setAccountConfirm && !isSetAccountLoading,
        loading: isSetAccountLoading,
      }}
    >
      <StyledForm
        form={form}
        layout="vertical"
        onFinish={(values) => handleSetAccount(values as unknown as FieldType)}
        onValuesChange={handleSetAccountFormChange}
        validateTrigger="onBlur"
        autoComplete="off"
      >
        <Form.Item<FieldType>
          label="Enter Account"
          name="account"
          rules={[
            { required: true, message: 'account is required' },
          ]}
        >
          <StyledInput placeholder="account" />
        </Form.Item>
        <Form.Item<FieldType>
          label="Enter Password"
          name="password"
          rules={[
            { required: true, message: 'password is required' },
          ]}
        >
          <PasswordInput placeholder="password" />
        </Form.Item>
      </StyledForm>
    </Modal>
  );
}

export default SetAccountModal;