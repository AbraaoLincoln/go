<?xml version="1.0"?>

<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">

<xsl:template match="/">
  <key><xsl:value-of select="Payload/Key"/></key>
  <xsl:text>&#10;</xsl:text>
  <xsl:for-each select="Payload/Properties">
    <Properties><xsl:value-of select="."/></Properties>
    <xsl:text>&#10;</xsl:text>
   </xsl:for-each>
</xsl:template>

</xsl:stylesheet> 