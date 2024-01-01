import React from 'react'
import { Image, Text, TouchableOpacity, View } from 'react-native'
import { Icons, colors } from '../../theme'
import styles from './AuthOptionScreen.style'
import { AuthSocialButton } from '../../components'



const AuthOptionScreen = () => {
    const { light, dark } = colors;
    let isDarkMode = false
    return (
        <View style={[{ backgroundColor: colors[isDarkMode ? 'dark' : 'light'].background }, styles.container]}>
            <View style={{ alignItems: 'center', marginBottom: 25 }}>
                <Text style={{ fontSize: 26, fontWeight: '500', color: isDarkMode ? dark.text : light.text }}>Let's you in</Text>
            </View>
            <AuthSocialButton title='Continue with Facebook' icon={Icons.Facebook} isDarkMode={isDarkMode} />
            <AuthSocialButton title='Continue with Google' icon={Icons.Google} isDarkMode={isDarkMode} />
        </View>
    )
}

export default AuthOptionScreen

