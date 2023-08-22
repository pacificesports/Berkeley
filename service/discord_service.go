package service

import (
	"berkeley/config"
	"berkeley/model"
	"berkeley/utils"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

var Discord *discordgo.Session

func ConnectDiscord() {
	dg, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		utils.SugarLogger.Errorln("Error creating Discord session, ", err)
		return
	}
	Discord = dg
	_, err = Discord.ChannelMessageSend(config.DiscordChannel, ":white_check_mark: "+config.Service.Name+" v"+config.Version+" online! `[ENV = "+config.Env+"]`")
	if err != nil {
		utils.SugarLogger.Errorln("Error sending Discord message, ", err)
		return
	}
}

func DiscordLogNewSchool(school model.School) {
	var embeds []*discordgo.MessageEmbed
	var fields []*discordgo.MessageEmbedField
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "ID",
		Value:  school.ID,
		Inline: false,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Type",
		Value:  school.Type,
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Verified",
		Value:  strconv.FormatBool(school.Verified),
		Inline: true,
	})
	embeds = append(embeds, &discordgo.MessageEmbed{
		Title: "New School Created!",
		Color: 62719,
		Author: &discordgo.MessageEmbedAuthor{
			URL:     school.Website,
			Name:    school.Name,
			IconURL: school.IconURL,
		},
		Fields: fields,
	})
	_, err := Discord.ChannelMessageSendEmbeds(config.DiscordChannel, embeds)
	if err != nil {
		utils.SugarLogger.Errorln(err.Error())
	}
}

func DiscordLogUpdatedSchool(school model.School) {
	var embeds []*discordgo.MessageEmbed
	var fields []*discordgo.MessageEmbedField
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "ID",
		Value:  school.ID,
		Inline: false,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Type",
		Value:  school.Type,
		Inline: true,
	})
	fields = append(fields, &discordgo.MessageEmbedField{
		Name:   "Verified",
		Value:  strconv.FormatBool(school.Verified),
		Inline: true,
	})
	embeds = append(embeds, &discordgo.MessageEmbed{
		Title: "School Updated!",
		Color: 11909047,
		Author: &discordgo.MessageEmbedAuthor{
			URL:     school.Website,
			Name:    school.Name,
			IconURL: school.IconURL,
		},
		Fields: fields,
	})
	_, err := Discord.ChannelMessageSendEmbeds(config.DiscordChannel, embeds)
	if err != nil {
		utils.SugarLogger.Errorln(err.Error())
	}
}
