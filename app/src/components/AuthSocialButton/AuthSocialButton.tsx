import React from 'react'
import { Image, StyleSheet, Text, TouchableOpacity, ImageSourcePropType } from 'react-native'
import { colors } from '../../theme'


interface Props {
    isDarkMode: boolean,
    icon: ImageSourcePropType,
    title: string
}

const AuthSocialButton = ({ isDarkMode, icon, title }: Props) => {
    const { light, dark } = colors;
    return (
        <TouchableOpacity style={[
            { backgroundColor: isDarkMode ? dark.primaryLight : light.white, borderWidth: isDarkMode ? undefined : 0.5 },
            styles.authButtonContainer]}
            activeOpacity={0.7}>
            <Image source={icon} style={styles.socialIcon} />
            <Text style={[{ color: isDarkMode ? dark.white : light.black }, styles.title]}>{title}</Text>
        </TouchableOpacity>
    )
}

export default AuthSocialButton

const styles = StyleSheet.create({
    authButtonContainer: {
        flexDirection: 'row',
        height: 60,
        justifyContent: 'center',
        alignItems: 'center',
        gap: 30,
        borderColor: colors.light.gray,
        marginVertical: 10,
        width: '90%',
        alignSelf: 'center',
        borderRadius: 12,
        paddingVertical: 6
    },
    socialIcon: {
        width: 30,
        height: 30
    },
    title: {
        fontSize: 18,
        fontWeight: '500',
    }
})