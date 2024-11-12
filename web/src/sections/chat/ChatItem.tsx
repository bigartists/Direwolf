import { memo } from 'react';

// import cuAvatar from 'assets/images/assistantAvator.png'
// import userAvartar from 'assets/images/chat-user.svg'

import cuAvatar from 'src/assets/images/v3/icon/taichu2.svg';
import userAvartar from 'src/assets/images/v3/icon/user.svg';

import Markdown from 'src/components/Markdown';

import { RoleTypes } from './type';
import { IChat } from './ModelClass/BaseChat';

export const ChatItem = memo(
  (props: {
    item: any;
    chatProps: IChat;
    onUserSubmit: (values: {
      text?: string;
      picture?: string;
      pictureType?: 'url' | 'base64';
      related?: any[];
    }) => Promise<void>;
  }) => {
    const { item, onUserSubmit, chatProps } = props;

    return (
      <div>
        <div>
          <img src={item.role === RoleTypes.user ? userAvartar : cuAvatar} alt="" />
        </div>
        <div>
          {item.role === RoleTypes.user ? (
            !item.picture ? (
              <div>{item.content}</div>
            ) : (
              <div>
                <div>
                  {item?.related ? (
                    <>
                      <div className="py-4 px-2">猜您想问：</div>
                      <div className="flex gap-[4px] flex-col px-2" />
                    </>
                  ) : null}
                </div>
                <div>{item.content}</div>
              </div>
            )
          ) : (
            <>
              {item.loading && (
                <div className="p-4 flex justify-center items-center">loading...</div>
              )}
              {item.type === 'error' ? (
                <div className="p-4">{item.content}</div>
              ) : (
                <div className="overflow-hidden w-full">
                  <Markdown content={item.content || ''} />
                </div>
              )}
            </>
          )}
        </div>
      </div>
    );
  }
);

export default ChatItem;
